
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RegistrationService } from '../../../services/Registration.service';
import { Registration } from '../../../models/Registration';

@Component({
    selector: 'app-index-registration',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRegistrationComponent implements OnInit {

    registrations: Registration[] = [];

    constructor(
        private router: Router,
        private service: RegistrationService
) {}

    ngOnInit(): void {
        this.getRegistrations();
}

    getRegistrations(): void {
        this.service.getRegistrations().subscribe((res) => {
        this.registrations = res;
    });
}

    deleteRegistration(id: any): void {
        this.service.deleteRegistration(id)
            .subscribe(() => {
                this.getRegistrations();
            });
    }
}