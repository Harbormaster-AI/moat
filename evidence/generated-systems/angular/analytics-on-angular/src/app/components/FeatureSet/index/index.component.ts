
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FeatureSetService } from '../../../services/FeatureSet.service';
import { FeatureSet } from '../../../models/FeatureSet';

@Component({
    selector: 'app-index-featureSet',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexFeatureSetComponent implements OnInit {

    featureSets: FeatureSet[] = [];

    constructor(
        private router: Router,
        private service: FeatureSetService
) {}

    ngOnInit(): void {
        this.getFeatureSets();
}

    getFeatureSets(): void {
        this.service.getFeatureSets().subscribe((res) => {
        this.featureSets = res;
    });
}

    deleteFeatureSet(id: any): void {
        this.service.deleteFeatureSet(id)
            .subscribe(() => {
                this.getFeatureSets();
            });
    }
}