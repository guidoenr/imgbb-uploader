# imgbb-uploader
Super light-weight linux-cli tool to upload images to the Imgbb free images storage in less than **2 seconds**!
- Maximum Size of image:`15MB`
- Maximum images p/hour: `20`

## Requirements
- You must set the `IMGBB_API_KEY` (as a local env) with your api key obtained from [ImgBB Free API](https://api.imgbb.com/). Example:
  ```bash
  export IMGBB_API_KEY=1234567thisIsAnInvalidKey
  ```
  Once you do that, you're ready 

## Instalation [auto]
- There is an `install.sh` file inside this repo, you can `sudo bash install.sh`

## Instalation [manual]
- Only move the binary in this repo to `/usr/local/bin/`. Example: 
    ```bash
    # upload is the binary file name
    sudo cp <yourPathTo>/imbgg-uploader/upload /usr/local/bin/upload
    ```
- If you don't trust, you can check the code by your own and do a `go build -o "binaryName"` and then move that to your bin path.


## Usage
- `upload exampleImage.png`
- `upload ../../anotherImage.jpg` (the tool is relative to your current working directory)

## Example Output
```json
{
  "data": {
    "thumb": {
      "filename": "p.png",
      "mime": "image/png",
      "url": "https://i.ibb.co/fr15bqj/p.png"
    },
    "delete_url": "https://ibb.co/fr15bqj/63869b09cb22ac411898d3470108ec39"
  },
  "success": true,
  "status": 200,
  "error": {}
}
```
Enjoy!


### @guidoenr TODO
- work with absolute path
- progress bar
- go routines



