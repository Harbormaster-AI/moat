import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MROFacilityService } from '../../../services/MROFacility.service';
import { MROFacility } from '../../../models/MROFacility';
import { SubBaseComponent } from '../../MROFacility/sub.base.component';

@Component({
    selector: 'app-create-mROFacility',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMROFacilityComponent extends SubBaseComponent implements OnInit {

    title = 'Add MROFacility';

    mROFacilityForm: FormGroup;
    mROFacility: MROFacility;

    constructor( http: HttpClient,
        private mROFacilityService: MROFacilityService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMROFacility(name, approvalScope, address, Appointments, WorkOrders): void {
        this.mROFacilityService
        .addMROFacility(name, approvalScope, address, Appointments, WorkOrders)
            .subscribe(() => {
                this.router.navigate(['/indexMROFacility']);
            });
    }

    ngOnInit(): void {
    }
}