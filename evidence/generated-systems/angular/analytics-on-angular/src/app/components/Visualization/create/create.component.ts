import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { VisualizationService } from '../../../services/Visualization.service';
import { Visualization } from '../../../models/Visualization';
import { SubBaseComponent } from '../../Visualization/sub.base.component';

@Component({
    selector: 'app-create-visualization',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateVisualizationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Visualization';

    visualizationForm: FormGroup;
    visualization: Visualization;

    constructor( http: HttpClient,
        private visualizationService: VisualizationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.visualizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      options: ['', Validators.required],
      Dashboard: ['', ],
      Report: ['', ],
      Metrics: ['', ],
      Dimensions: ['', ],
      Datasets: ['', ],
      ChartType: ['', ]
        });
    }

    
    addVisualization(title, options, Dashboard, Report, Metrics, Dimensions, Datasets, ChartType): void {
        this.visualizationService
        .addVisualization(title, options, Dashboard, Report, Metrics, Dimensions, Datasets, ChartType)
            .subscribe(() => {
                this.router.navigate(['/indexVisualization']);
            });
    }

    ngOnInit(): void {
    }
}