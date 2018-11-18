from __future__ import print_function
from __future__ import unicode_literals
from __future__ import division
from __future__ import absolute_import
from builtins import str
from future import standard_library
standard_library.install_aliases()
import os
import sys
import subprocess
import platform

globalMainDatabaseName="sandbox"
globalTestDatabaseName="sandboxtest"

def execute(param, **kwargs):
	try:
		subprocess.check_call(param, **kwargs)
	except subprocess.CalledProcessError as err:
		sys.exit(err.returncode)

def buildserver( ):
	go_get = ["go", "get"]
	execute( go_get + ["github.com/go-sql-driver/mysql"])
	execute( go_get + ["github.com/lib/pq"])
	execute( go_get + ["github.com/lib/pq"])
	if os.path.exists("server"):
		os.chdir("server")
	else:
		sys.exit(1)
	execute(["go", "build"])
	os.chdir("..")

def testserver():
	dbpass = ""
	dbuser = ""
	dbport = 5432
	if ("PGPASSWORD" in os.environ ) and ("PGUSER" in os.environ):
		dbpass = os.environ["PGPASSWORD"]
		dbuser = os.environ["PGUSER"]
	else:
		print("Please set environment variables PGPASSWORD andn PGUSER")
		sys.exit(1)
	if ("PGPORT" in os.environ):
		dbport = os.environ["PGPORT"]
	os.putenv("GOCACHE", "off")
	os.putenv("acs_database", "postgres")
	os.putenv("acs_dbuser", dbuser)
	os.putenv("acs_dbpass", dbpass)
	os.putenv("acs_onlineDbName", globalMainDatabaseName)
	os.putenv("acs_onlineDbNameTest", globalTestDatabaseName)
	os.putenv("acs_hostname", "127.0.0.1")
	os.putenv("acs_hostport", str(dbport))
	if os.path.exists("server"):
		os.chdir("server")
	else:
		sys.exit(1)
	execute(["go", "test", os.path.join(".","...")])
	os.chdir("..")
	os.putenv("GOCACHE", "on")
	
def buildclient():
	if platform.system() == "Windows":
		execute(["nuget", "restore", os.path.join("client","TimeTracker","TimeTracker","packages.config"), "-OutputDirectory", os.path.join("client","TimeTracker","packages")])
		execute(["msbuild", os.path.join("client","TimeTracker","TimeTracker.sln")])

def testclient():
	if platform.system() == "Windows":
		execute(["nunit3-console", "--where", "\"namespace==TimeTracker\"", os.path.join("bin","Debug","TimeTracker.exe")])

def rebuildDatabase():
	dbpass=""
	dbuser=""
	if ("PGPASSWORD" in os.environ ) and ("PGUSER" in os.environ):
		dbpass = os.environ["PGPASSWORD"]
		dbuser = os.environ["PGUSER"]
	else:
		print("Please set environment variables PGPASSWORD and PGUSER")
		sys.exit(1)
	subprocess.call(["createdb", globalMainDatabaseName])
	subprocess.call(["createdb", globalTestDatabaseName])
	execute(["psql", "-U", dbuser, "-d", globalMainDatabaseName, "-f", "sandbox.sql"])
	execute(["psql", "-U", dbuser, "-d", globalTestDatabaseName, "-f", "sandbox.sql"])

def main():
	if len(sys.argv) < 2:
		print("please specify command: builddb, buildserver, testserver, buildclient")
		sys.exit(1)
	if sys.argv[1] == "builddb":
		rebuildDatabase()
	elif sys.argv[1] == "buildserver":
		buildserver()
	elif sys.argv[1] == "testserver":
		testserver()
	elif sys.argv[1] == "buildclient":
		buildclient()

if __name__ == "__main__":
    main()
