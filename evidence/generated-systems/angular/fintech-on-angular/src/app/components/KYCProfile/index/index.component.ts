
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { KYCProfileService } from '../../../services/KYCProfile.service';
import { KYCProfile } from '../../../models/KYCProfile';

@Component({
    selector: 'app-index-kYCProfile',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexKYCProfileComponent implements OnInit {

    kYCProfiles: KYCProfile[] = [];

    constructor(
        private router: Router,
        private service: KYCProfileService
) {}

    ngOnInit(): void {
        this.getKYCProfiles();
}

    getKYCProfiles(): void {
        this.service.getKYCProfiles().subscribe((res) => {
        this.kYCProfiles = res;
    });
}

    deleteKYCProfile(id: any): void {
        this.service.deleteKYCProfile(id)
            .subscribe(() => {
                this.getKYCProfiles();
            });
    }
}