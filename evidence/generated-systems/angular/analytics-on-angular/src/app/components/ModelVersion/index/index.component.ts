
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ModelVersionService } from '../../../services/ModelVersion.service';
import { ModelVersion } from '../../../models/ModelVersion';

@Component({
    selector: 'app-index-modelVersion',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexModelVersionComponent implements OnInit {

    modelVersions: ModelVersion[] = [];

    constructor(
        private router: Router,
        private service: ModelVersionService
) {}

    ngOnInit(): void {
        this.getModelVersions();
}

    getModelVersions(): void {
        this.service.getModelVersions().subscribe((res) => {
        this.modelVersions = res;
    });
}

    deleteModelVersion(id: any): void {
        this.service.deleteModelVersion(id)
            .subscribe(() => {
                this.getModelVersions();
            });
    }
}