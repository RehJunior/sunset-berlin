To install the `sunset-berlin` DEB package from https://apt.fury.io/cyberants/ on Ubuntu:

1. Add the repository:
   ```
   echo "deb [trusted=yes] https://apt.fury.io/cyberants/ /" |sudo tee /etc/apt/sources.list.d/fury.list
   ```

2. Update the package list:
   ```
   sudo apt update
   ```

3. Install the package:
   ```
   sudo apt install sunset-berlin
   ```
