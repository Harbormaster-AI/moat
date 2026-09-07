
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FeatureService } from '../../../services/Feature.service';
import { Feature } from '../../../models/Feature';

@Component({
    selector: 'app-index-feature',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFeatureComponent implements OnInit {

    features: Feature[] = [];

    constructor(
        private router: Router,
        private service: FeatureService
) {}

    ngOnInit(): void {
        this.getFeatures();
}

    getFeatures(): void {
        this.service.getFeatures().subscribe((res) => {
        this.features = res;
    });
}

    deleteFeature(id: any): void {
        this.service.deleteFeature(id)
            .subscribe(() => {
                this.getFeatures();
            });
    }
}