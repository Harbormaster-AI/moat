
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { APIClientService } from '../../../services/APIClient.service';
import { APIClient } from '../../../models/APIClient';

@Component({
    selector: 'app-index-aPIClient',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAPIClientComponent implements OnInit {

    aPIClients: APIClient[] = [];

    constructor(
        private router: Router,
        private service: APIClientService
) {}

    ngOnInit(): void {
        this.getAPIClients();
}

    getAPIClients(): void {
        this.service.getAPIClients().subscribe((res) => {
        this.aPIClients = res;
    });
}

    deleteAPIClient(id: any): void {
        this.service.deleteAPIClient(id)
            .subscribe(() => {
                this.getAPIClients();
            });
    }
}