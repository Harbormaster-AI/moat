import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ServiceBulletinService } from '../../../services/ServiceBulletin.service';
import { ServiceBulletin } from '../../../models/ServiceBulletin';
import { SubBaseComponent } from '../../ServiceBulletin/sub.base.component';

@Component({
    selector: 'app-create-serviceBulletin',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateServiceBulletinComponent extends SubBaseComponent implements OnInit {

    title = 'Add ServiceBulletin';

    serviceBulletinForm: FormGroup;
    serviceBulletin: ServiceBulletin;

    constructor( http: HttpClient,
        private serviceBulletinService: ServiceBulletinService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.serviceBulletinForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  bulletinNumber: ['', Validators.required],
      WorkOrders: ['', ],
      Variants: ['', ],
      Category: ['', ]
        });
    }

    
    addServiceBulletin(bulletinNumber, WorkOrders, Variants, Category): void {
        this.serviceBulletinService
        .addServiceBulletin(bulletinNumber, WorkOrders, Variants, Category)
            .subscribe(() => {
                this.router.navigate(['/indexServiceBulletin']);
            });
    }

    ngOnInit(): void {
    }
}