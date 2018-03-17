# Lightroom <--> External Drive <--> S3 tool

[https://github.com/byronmansfield/lrbkup](https://github.com/byronmansfield/lrbkup)

This is a tool designed to help me manage my Lightroom photos. It ensures data integrity by backing up to an external drive(s) and to S3. It also helps manage my local drives disk space.

I wrote this to solve some problems I had while on the road and suffered a hard drive crash resulting in a loss of my photographs. In addition to having limited hard drive space and limited internet connection.

## How it works

### Import to Lightroom

So after taking some photos and they only live on your SD card, you go to import them into your computers drive via Lightroom. While yes you can skip installing them directly to your local hard drive and tell LR to install them somewhere else, like an external hard drive. It is a bit slower to load them and edit them. So I usually like to start local and move them off of my local if I am not going to be using them within Lightroom anytime soon.

### Back up your local to an external hard drive

So while having them local is great for now, you have 2 potential issues.

If you have imported them on to your local and wiped your SD cards to take more photos (like ya do) then that means your only copy of the photos lives on your computers local hard drive. What happens if it is lost, stolen, or crashes. *POOF* They are gone.

Second of all, eventually you will fill up your computers hard drive. So copying them to an external hard drive solves both of these problems. You don't actually have to move your photos within LR to the external hard drive yet. We just want to keep a second copy in case of a dooms day experience.

### Move your photos in Lightroom to the external drive

When you are ready to move them out of your local for extra space, you will have to do this by the native LR UI. It's easy. Just select the photos and tell it to update it's location to the backed up drive.

Once this is done, then you are free to clean up your local hard drive if you want. Though if you have not backed up your external drive to the cloud then you will again be at risk. Your external drive is the only place your photos live. So if it crashes (like in my case) you will loose those photos.

### Copy the photos to some cloud storage

Right now I have only integrated with S3, but in the future I may add other cloud storage platforms to the tool.

The Lightroom Backup tool can sync your drive with S3 then you will always have a copy unless S3 goes down. Which is not likely to happen.

### Now we free up some local disk space

With our images on our local drive, external drive, and in the cloud, we can safely remove them from one of the 3 locations and have 2 copies. In case our external drive bites the dust, we can always pull them down from S3.

The LR backup tool is designed to help you with housekeeping on your local drive, it even has a reporting tool that will show you which ones have backed up copies in an external and S3. This way we don't make any mistakes of removing some photos before they get backed up. It also gives you a report about how much disk space was cleaned up.

## Features

 - A reporting tool showing you where your photos are backed up
 - Local disk manager to assist with freeing up space
 - Sync new disk from S3 in case of failed external disk
