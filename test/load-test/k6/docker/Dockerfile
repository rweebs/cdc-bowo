# Use the loadimpact/k6 image as the base image
FROM loadimpact/k6

# Copy your k6 script to the container
COPY script.js .

# Set the default command to run your k6 script
CMD ["run", "script.js"]