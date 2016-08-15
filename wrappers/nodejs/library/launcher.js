/**
 * Launcher module. A useful module to help launching platform dependent binaries.
 * This module deals with the overhead of supported operating systems, architectures and the variety of mechanisms
 * used to launch applications on different platforms.
 * @module wpwithin/launcher
 */
module.exports = {
  Config : Config,
  launch : launch
};

/**
 * Launcher config.
 * @class Config
 */
function Config(requiredOSs, requiredArchs) {

  // Arch can be { 'arm', 'ia32', or 'x64' } as per https://nodejs.org/dist/latest-v5.x/docs/api/process.html#process_process_arch
  // OS can be { 'darwin', 'freebsd', 'linux', 'sunos' or 'win32' } as per https://nodejs.org/dist/latest-v5.x/docs/api/process.html#process_process_platform

  this.requiredOSs = requiredOSs;
  this.requiredArchs = requiredArchs;
}

/**
  * Launch an application
  * config
  * path is the path to the executable
  * flags are optional but a way so specify extra information
  */
function launch(config, path, flags, callback) {

  // Determine the OS and Architecture this application is currently running on
  var hostOS = detectHostOS().toLowerCase();
  var hostArchitecture = detectHostArchitecture().toLowerCase();


  if(validateConfig(config, hostOS, hostArchitecture)) {

    switch(hostOS) {

      case "darwin":
      launchDarwin(path, flags, callback);
      break;

      case "linux":
      launchLinux(path, flags, callback);
      break;

      case "win32":
      launchWindows(path, flags, callback);
      break;

      default:
      throw new Error("Unable to launch binary on host architecture (Unsupported by launcher)(Host=%s)", hostOS);
      break;
    }
  } else {

      console.log("Invalid OS/Architecture combination detected");
  }
}

/**
 * Determine what Operating System this application is running on
 */
function detectHostOS() {

  return process.platform;
}

/**
  * Detect the CPU Atchitecture that this application is currently running on
  **/
function detectHostArchitecture() {

  return process.arch;
}

function launchDarwin(path, flags, callback) {

  var util = require('util');

  console.log("launching Darwin application");

  var exec = require('child_process').exec;
  var cmd = util.format('%s %s', path, flags == null ? "" : flags);

  exec(cmd, function(error, stdout, stderr) {

    return callback(error, stdout, stderr);
  });
}

function launchLinux(path, flags, callback) {

  console.log("launching Linux application");
}

function launchWindows(path, flags, callback) {

  console.log("launching Windows application");
}
/**
 * For a given config, determine if it matches the specified hostOS and hostArchitecture
 */
function validateConfig(config, hostOS, hostArchitecture) {

  // Validate detected parameters against config
  var validOS = false;
  var validArch = false;

  for(i in config.requiredOSs) {

    if(config.requiredOSs[i].toLowerCase() === hostOS) {

      validOS = true;
    }
  }

  for(i in config.requiredArchs) {

    if(config.requiredArchs[i].toLowerCase() === hostArchitecture) {

      validArch = true;
    }
  }

  return validOS && validArch;
}
