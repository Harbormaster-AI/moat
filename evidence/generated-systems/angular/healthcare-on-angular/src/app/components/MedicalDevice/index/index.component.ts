
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MedicalDeviceService } from '../../../services/MedicalDevice.service';
import { MedicalDevice } from '../../../models/MedicalDevice';

@Component({
    selector: 'app-index-medicalDevice',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMedicalDeviceComponent implements OnInit {

    medicalDevices: MedicalDevice[] = [];

    constructor(
        private router: Router,
        private service: MedicalDeviceService
) {}

    ngOnInit(): void {
        this.getMedicalDevices();
}

    getMedicalDevices(): void {
        this.service.getMedicalDevices().subscribe((res) => {
        this.medicalDevices = res;
    });
}

    deleteMedicalDevice(id: any): void {
        this.service.deleteMedicalDevice(id)
            .subscribe(() => {
                this.getMedicalDevices();
            });
    }
}