
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TargetingProfileService } from '../../../services/TargetingProfile.service';
import { TargetingProfile } from '../../../models/TargetingProfile';

@Component({
    selector: 'app-index-targetingProfile',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexTargetingProfileComponent implements OnInit {

    targetingProfiles: TargetingProfile[] = [];

    constructor(
        private router: Router,
        private service: TargetingProfileService
) {}

    ngOnInit(): void {
        this.getTargetingProfiles();
}

    getTargetingProfiles(): void {
        this.service.getTargetingProfiles().subscribe((res) => {
        this.targetingProfiles = res;
    });
}

    deleteTargetingProfile(id: any): void {
        this.service.deleteTargetingProfile(id)
            .subscribe(() => {
                this.getTargetingProfiles();
            });
    }
}