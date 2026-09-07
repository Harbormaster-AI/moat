
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SubscriberService } from '../../../services/Subscriber.service';
import { Subscriber } from '../../../models/Subscriber';

@Component({
    selector: 'app-index-subscriber',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexSubscriberComponent implements OnInit {

    subscribers: Subscriber[] = [];

    constructor(
        private router: Router,
        private service: SubscriberService
) {}

    ngOnInit(): void {
        this.getSubscribers();
}

    getSubscribers(): void {
        this.service.getSubscribers().subscribe((res) => {
        this.subscribers = res;
    });
}

    deleteSubscriber(id: any): void {
        this.service.deleteSubscriber(id)
            .subscribe(() => {
                this.getSubscribers();
            });
    }
}