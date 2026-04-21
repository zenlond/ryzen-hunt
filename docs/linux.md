
## 🚀 Setup Instructions

### 1. Download Server and Configure Hosts

1. Go to [Releases](https://github.com/zenlond/ryzen-hunt/releases/tag/1.0)
2. Download the latest `server_linux.zip` and extract it to a folder.

Edit `hosts.txt` with your PC's IPv4 address:
- Find out your computer's **IPv4 address**
- Replace `192.168.X.X` in `hosts.txt` with this address
- ⚠️ Make sure to use the **Wi-Fi network IPv4 address** (not Ethernet, VPN, VirtualBox, etc.)

###

### 2. 📱 Prepare Android Device

1. Install **Ryzen Hunt Event APK**
2. Install **Hosts GO** app
3. Set date to **August 2017**
4. Transfer `hosts.txt` to your phone

Both devices must connect to the same Wi-Fi network (avoid split bands like 5GHz).

### 3. ⚙️ Configure Hosts in Hosts GO

1. Open **Hosts Editor**
2. Tap **⋮** → **Import hosts**
3. Select `hosts.txt`
4. Enable the **Hosts change switch**

### 4. ▶️ Start Server and App

1. Make the server executable: `chmod +x server`
2. Run the server: `sudo ./server`
3. Launch the Ryzen Hunt App on your Android device
4. If `server` displays "ok" and the app shows the country selection screen, the setup is complete and working correctly

### 5. 📝 Register Account

1. Choose **Singapore**
2. Register with any credentials (no validation required):
    - Username: `anon`
    - Email: `anon@email.com`

### 6. 📸 Access AR Features

Tap **Hunt**, point camera at promotional poster, and the AR animation will display. Enjoy
