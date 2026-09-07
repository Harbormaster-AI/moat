
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HealthSystemService } from '../../../services/HealthSystem.service';
import { HealthSystem } from '../../../models/HealthSystem';

@Component({
    selector: 'app-index-healthSystem',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexHealthSystemComponent implements OnInit {

    healthSystems: HealthSystem[] = [];

    constructor(
        private router: Router,
        private service: HealthSystemService
) {}

    ngOnInit(): void {
        this.getHealthSystems();
}

    getHealthSystems(): void {
        this.service.getHealthSystems().subscribe((res) => {
        this.healthSystems = res;
    });
}

    deleteHealthSystem(id: any): void {
        this.service.deleteHealthSystem(id)
            .subscribe(() => {
                this.getHealthSystems();
            });
    }
}