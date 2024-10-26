# wtf-deployment

This repository contains required files to deploy what-da-flac application using docker compose.

## Requirements

* [Docker](https://www.docker.com/products/docker-desktop/)
* [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

Update packages

```
sudo apt update
sudo apt install -y unzip build-essential
```

Generate ssh key and make sure it is added to Github deploy keys as read-only.

Installation and configuration

```
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker $USER
newgrp docker

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
rm -rf aws

git config --global url."git@github.com:".insteadOf "https://github.com/"
echo 'export GIT_SSH_COMMAND="ssh -o BatchMode=yes -o StrictHostKeyChecking=no"' >> ~/.bashrc
```

Source environment and clone repository

```
source ~/.bashrc
git clone https://github.com/tech-component/wtf-deployment
```

## Backend

### Requisites

Copy the sample files and modify `.env.credentials` accordingly

```
cp sample.env.credentials .env.credentials
```

## Services

### Registration

```
sudo cp ubuntu/wtf.service /etc/systemd/system/wtf.service
sudo systemctl enable wtf.service
sudo service start wtf
```

### Start

First, make sure your code is up to date with latest changes. Run this command to pull changes from Github:

```
make pull
```

Next, start backend services:

```
make start
```

You can run all these commands together in one line too:

```
make pull build-all start
```

### Stop

Stop services with this command:

```
make stop
```

## SSL Certificates

1. Install Certbot and NGINX Plugin

```
sudo apt update
sudo apt install -y certbot python3-certbot-nginx nginx
sudo certbot --nginx
```

Once installed, restart NGINX service

```
sudo systemctl restart nginx
```

2. Enable Auto-Renewal

```
sudo certbot renew --dry-run
```

## NGINX Permissions

```
sudo chown -R ubuntu:www-data /home/ubuntu/ui
sudo chmod -R 770 /home/ubuntu/ui
sudo chown -R ubuntu:www-data /home/ubuntu/ui/build
sudo chmod -R 750 /home/ubuntu/ui/build
sudo chmod 755 /home/ubuntu
sudo chmod 755 /home/ubuntu/ui
sudo chmod 750 /home/ubuntu/ui/build
sudo chmod 644 /home/ubuntu/ui/build/index.html
```

Restart service

```
sudo service nginx restart
```

Check

```
ls -ld /home/ubuntu/ui
```


## SSO

We use SSO from Google for the time being. You will need a Google account in order to run the applications.

[Here](https://console.cloud.google.com/apis/credentials?referrer=search&project=openid-test-397315) is the right spot

![google cloud platform screenshot](./assets/google-cloud-console-credentials.png)

These are the right settings:

![google cloud platform screenshot](./assets/gcp-sso-configuration-web-app.png)

PublicitUX system automatically creates users, once you login with Google SSO.

## AWS

## Access Keys

Create access keys for user `wtf-app`

```
aws --profile wtf-mauleyzaola iam create-access-key --user-name wtf-app
```

## Configure AWS Profile

```
aws --profile wtf configure
```

## TODO

Repositories

```
docker pull 160885250498.dkr.ecr.us-east-2.amazonaws.com/wtf-user-service
```