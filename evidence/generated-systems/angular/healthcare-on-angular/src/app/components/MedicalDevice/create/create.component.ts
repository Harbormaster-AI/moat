import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MedicalDeviceService } from '../../../services/MedicalDevice.service';
import { MedicalDevice } from '../../../models/MedicalDevice';
import { SubBaseComponent } from '../../MedicalDevice/sub.base.component';

@Component({
    selector: 'app-create-medicalDevice',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMedicalDeviceComponent extends SubBaseComponent implements OnInit {

    title = 'Add MedicalDevice';

    medicalDeviceForm: FormGroup;
    medicalDevice: MedicalDevice;

    constructor( http: HttpClient,
        private medicalDeviceService: MedicalDeviceService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMedicalDevice(udi, manufacturer, Patient, Observations, SoftwareUpdates, DeviceType, ConnectivityStatus): void {
        this.medicalDeviceService
        .addMedicalDevice(udi, manufacturer, Patient, Observations, SoftwareUpdates, DeviceType, ConnectivityStatus)
            .subscribe(() => {
                this.router.navigate(['/indexMedicalDevice']);
            });
    }

    ngOnInit(): void {
    }
}