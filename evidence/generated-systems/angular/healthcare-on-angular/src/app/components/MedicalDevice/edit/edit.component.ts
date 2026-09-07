import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MedicalDeviceService } from '../../../services/MedicalDevice.service';
import { SubBaseComponent } from '../../MedicalDevice/sub.base.component';


@Component({
    selector: 'app-edit-medicalDevice',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMedicalDeviceComponent extends SubBaseComponent implements OnInit {

    title = 'Edit MedicalDevice';

    medicalDeviceForm: FormGroup;
    medicalDevice: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MedicalDeviceService,
        private fb: FormBuilder
) {
        super(http);
        this.medicalDeviceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  udi: ['', Validators.required],
      manufacturer: ['', Validators.required],
      Patient: ['', ],
      Observations: ['', ],
      SoftwareUpdates: ['', ],
      DeviceType: ['', ],
      ConnectivityStatus: ['', ]
        });
    }

    
    updateMedicalDevice(udi, manufacturer, Patient, Observations, SoftwareUpdates, DeviceType, ConnectivityStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMedicalDevice(udi, manufacturer, Patient, Observations, SoftwareUpdates, DeviceType, ConnectivityStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMedicalDevice']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMedicalDevice(params['id']).subscribe(res => {
                this.medicalDevice = res;
            });
        });
    }
}