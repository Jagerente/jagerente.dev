---
title: "Vestan - VR Media Streaming Platform"
description: "Software Engineer"
url: ""
type: "work"
references:
  - title: "vestan.tech"
    url: "https://vestan.tech/"
  - title: "LinkedIn"
    url: "https://www.linkedin.com/company/vestan-tech"
---

![Vestan](/images/projects/vestan.webp)

I currently work as a Software Engineer at a high-load platform specializing in media streaming.

I work on backend, architecture, and code reviews. I mainly use Go; plus PHP (Symfony) and Vue 3.

## Media Streaming & Architecture

Streaming is the core of our platform. When I joined, the service was unstable: we had race conditions, frequent bugs, and performance issues.

I refactored a large part of the codebase, fixed concurrency problems, and updated parts of the tech stack. After that, the system became much more stable, and the code is now cleaner and easier to maintain.

Another important part of my work is making the streaming service ready for scaling. I’m implementing cluster support so we can run multiple media servers, replicate capacity when needed, and route traffic to the best server. This improves content delivery for viewers and makes stream ingestion more reliable for creators. It also gives the platform a clear path to grow without rewriting the whole system later.

## New systems

I have designed and built multiple microservices for the platform.

The biggest and most complex one is a Private Messaging Service. It is built for high load and real-time communication. I implemented it from scratch and focused on reliability, scalability, and maintainability. It required careful work with data consistency, performance, and real-time delivery, because messaging systems break quickly under load if the design is weak.

## DX & processes

I strongly believe that a messy dev environment kills productivity.

Before, local setup was a "zoo" of configs, undocumented steps and third-party dependencies. It was hard to run the full ecosystem of services, and each developer had a different setup.

I worked on standardizing our workflow and code culture. I wrote clear documentation, improved tooling, and created a simple way to run multiple interacting services locally. Now, starting the whole project stack is a standard process, and developers can focus on building features instead of struggling with the environment.
