
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ServiceProviderService } from '../../../services/ServiceProvider.service';
import { ServiceProvider } from '../../../models/ServiceProvider';

@Component({
    selector: 'app-index-serviceProvider',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexServiceProviderComponent implements OnInit {

    serviceProviders: ServiceProvider[] = [];

    constructor(
        private router: Router,
        private service: ServiceProviderService
) {}

    ngOnInit(): void {
        this.getServiceProviders();
}

    getServiceProviders(): void {
        this.service.getServiceProviders().subscribe((res) => {
        this.serviceProviders = res;
    });
}

    deleteServiceProvider(id: any): void {
        this.service.deleteServiceProvider(id)
            .subscribe(() => {
                this.getServiceProviders();
            });
    }
}