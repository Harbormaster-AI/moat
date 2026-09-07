
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ServiceBulletinService } from '../../../services/ServiceBulletin.service';
import { ServiceBulletin } from '../../../models/ServiceBulletin';

@Component({
    selector: 'app-index-serviceBulletin',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexServiceBulletinComponent implements OnInit {

    serviceBulletins: ServiceBulletin[] = [];

    constructor(
        private router: Router,
        private service: ServiceBulletinService
) {}

    ngOnInit(): void {
        this.getServiceBulletins();
}

    getServiceBulletins(): void {
        this.service.getServiceBulletins().subscribe((res) => {
        this.serviceBulletins = res;
    });
}

    deleteServiceBulletin(id: any): void {
        this.service.deleteServiceBulletin(id)
            .subscribe(() => {
                this.getServiceBulletins();
            });
    }
}