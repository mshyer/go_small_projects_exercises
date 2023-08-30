// const process = require("node:process")
const fs = require('fs')
const path = process.cwd()

function findDirectories(path) {
  console.log(path);
  fs.readdir(path, (err, files) => {
    if (err) {
      
      console.error("error reading directory", err);
      return;
    }

    for (let file of files) {
      fs.stat(`${path}/${file}`, (err, stats) => {
        if (err) {
          console.error(err)
          return;
        }
        if (stats.isDirectory()) {
          findDirectories(`${path}/${file}`)
        } else {
          console.log(file);
        }
      });
    }
  });

}

module.exports = {
  findDirectories
}

findDirectories(path)