
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ContactService } from '../../../services/Contact.service';
import { Contact } from '../../../models/Contact';

@Component({
    selector: 'app-index-contact',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexContactComponent implements OnInit {

    contacts: Contact[] = [];

    constructor(
        private router: Router,
        private service: ContactService
) {}

    ngOnInit(): void {
        this.getContacts();
}

    getContacts(): void {
        this.service.getContacts().subscribe((res) => {
        this.contacts = res;
    });
}

    deleteContact(id: any): void {
        this.service.deleteContact(id)
            .subscribe(() => {
                this.getContacts();
            });
    }
}