
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { VisualizationService } from '../../../services/Visualization.service';
import { Visualization } from '../../../models/Visualization';

@Component({
    selector: 'app-index-visualization',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexVisualizationComponent implements OnInit {

    visualizations: Visualization[] = [];

    constructor(
        private router: Router,
        private service: VisualizationService
) {}

    ngOnInit(): void {
        this.getVisualizations();
}

    getVisualizations(): void {
        this.service.getVisualizations().subscribe((res) => {
        this.visualizations = res;
    });
}

    deleteVisualization(id: any): void {
        this.service.deleteVisualization(id)
            .subscribe(() => {
                this.getVisualizations();
            });
    }
}