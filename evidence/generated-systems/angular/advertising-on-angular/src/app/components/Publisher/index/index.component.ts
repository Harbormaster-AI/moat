
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PublisherService } from '../../../services/Publisher.service';
import { Publisher } from '../../../models/Publisher';

@Component({
    selector: 'app-index-publisher',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexPublisherComponent implements OnInit {

    publishers: Publisher[] = [];

    constructor(
        private router: Router,
        private service: PublisherService
) {}

    ngOnInit(): void {
        this.getPublishers();
}

    getPublishers(): void {
        this.service.getPublishers().subscribe((res) => {
        this.publishers = res;
    });
}

    deletePublisher(id: any): void {
        this.service.deletePublisher(id)
            .subscribe(() => {
                this.getPublishers();
            });
    }
}