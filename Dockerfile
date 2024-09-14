# Use ROS 2 Humble desktop image as base
FROM osrf/ros:humble-desktop-full

# Set the environment variable for the display and domain ID
ENV DISPLAY=:0
ENV ROS_DOMAIN_ID=37


### Install necessary packages for ROS and GUI applications
RUN apt-get update && apt-get upgrade -y
# ROS Tools
RUN apt-get update \
	&& apt-get install -y ros-humble-desktop \
	ros-dev-tools \
	ros-humble-rviz2 \
	ros-humble-xacro \
	ros-humble-turtlebot3* \
    ros-humble-navigation2 \
	ros-humble-nav2-bringup \
    ros-humble-ros2-control \
	ros-humble-ros2-controllers \
	ros-humble-gazebo-ros2-control \
	ros-humble-tf-transformations

# Other Tools
RUN apt-get update \
	&& apt-get install -y apt-utils \
	clang \
	libc6-dev \
	vim \
	git \
	nmap \
	sudo \
    screen \
	ranger \
	x11-apps \
	net-tools \
    libserial-dev \
	iputils-ping \
	python3-pip \
	python3-smbus

# Install Rust
# RUN curl https://sh.rustup.rs -sSf | sh -s -- -y

# Ensure Rust binaries are available in the PATH
# ENV PATH="/root/.cargo/bin:${PATH}"

# Install Go Lang
RUN cd /root/ \
	&& wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz \
	&& rm -rf /usr/local/go /usr/bin/go \
	&& tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz \
	&& export PATH=$PATH:/usr/local/go/bin

# go get github.com/tiiuae/rclgo/cmd/rclgo-gen/cmd@v0.0.0-20240131135202-56b24e11219b
	
# Clean up (For release image)
# RUN rm -rf /var/lib/apt/lists/* \

RUN pip3 install pyserial transforms3d


### Set up user
ARG USERNAME=ros
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create a non-root user 
RUN groupadd --gid $USER_GID $USERNAME \
	&& useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
	&& mkdir /home/$USERNAME/.config && chown $USER_UID:$USER_GID /home/$USERNAME/.config

# Set up sudo 
RUN echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
	&& chmod 0440 /etc/sudoers.d/$USERNAME

# Add new user to dialout 
RUN usermod -aG dialout ${USERNAME}


### Ros setup
USER ros

# Setup working directory
WORKDIR /home/$USERNAME/dev_ws

# Copy ROS packages from your local 'src' to the container
# COPY dev_ws/src /home/$USERNAME/dev_ws/src

RUN cd /home/$USERNAME/dev_ws \
	&& source /opt/ros/humble/setup.bash \
	&& colcon build
RUN echo "source /home/$USERNAME/dev_ws/install/setup.bash" >> /home/$USERNAME/.bashrc

RUN echo "export PATH=$PATH:/usr/local/go/bin" >> /home/$USERNAME/.bashrc
RUN echo "export LIBCLANG_PATH=/usr/lib/llvm-12/lib" >> /home/$USERNAME/.bashrc
RUN echo "source /opt/ros/humble/setup.bash" >> /home/$USERNAME/.bashrc
RUN echo "source /usr/share/colcon_argcomplete/hook/colcon-argcomplete.bash" >> /home/$USERNAME/.bashrc
RUN echo "set -o vi" >> /home/ros/.bashrc


USER root
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
CMD ["/bin/bash"]



