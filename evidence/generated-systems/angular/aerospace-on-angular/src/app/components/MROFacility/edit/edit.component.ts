import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MROFacilityService } from '../../../services/MROFacility.service';
import { SubBaseComponent } from '../../MROFacility/sub.base.component';


@Component({
    selector: 'app-edit-mROFacility',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMROFacilityComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MROFacility';

    mROFacilityForm: FormGroup;
    mROFacility: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MROFacilityService,
        private fb: FormBuilder
) {
        super(http);
        this.mROFacilityForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      approvalScope: ['', Validators.required],
      address: ['', Validators.required],
      Appointments: ['', ],
      WorkOrders: ['', ]
        });
    }

    
    updateMROFacility(name, approvalScope, address, Appointments, WorkOrders): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMROFacility(name, approvalScope, address, Appointments, WorkOrders, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMROFacility']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMROFacility(params['id']).subscribe(res => {
                this.mROFacility = res;
            });
        });
    }
}