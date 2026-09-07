
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InferenceEndpointService } from '../../../services/InferenceEndpoint.service';
import { InferenceEndpoint } from '../../../models/InferenceEndpoint';

@Component({
    selector: 'app-index-inferenceEndpoint',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInferenceEndpointComponent implements OnInit {

    inferenceEndpoints: InferenceEndpoint[] = [];

    constructor(
        private router: Router,
        private service: InferenceEndpointService
) {}

    ngOnInit(): void {
        this.getInferenceEndpoints();
}

    getInferenceEndpoints(): void {
        this.service.getInferenceEndpoints().subscribe((res) => {
        this.inferenceEndpoints = res;
    });
}

    deleteInferenceEndpoint(id: any): void {
        this.service.deleteInferenceEndpoint(id)
            .subscribe(() => {
                this.getInferenceEndpoints();
            });
    }
}