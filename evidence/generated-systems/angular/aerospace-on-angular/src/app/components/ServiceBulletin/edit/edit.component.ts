import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ServiceBulletinService } from '../../../services/ServiceBulletin.service';
import { SubBaseComponent } from '../../ServiceBulletin/sub.base.component';


@Component({
    selector: 'app-edit-serviceBulletin',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditServiceBulletinComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ServiceBulletin';

    serviceBulletinForm: FormGroup;
    serviceBulletin: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ServiceBulletinService,
        private fb: FormBuilder
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

    
    updateServiceBulletin(bulletinNumber, WorkOrders, Variants, Category): void {
        this.route.params.subscribe((params) => {

                        this.service.updateServiceBulletin(bulletinNumber, WorkOrders, Variants, Category, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexServiceBulletin']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getServiceBulletin(params['id']).subscribe(res => {
                this.serviceBulletin = res;
            });
        });
    }
}