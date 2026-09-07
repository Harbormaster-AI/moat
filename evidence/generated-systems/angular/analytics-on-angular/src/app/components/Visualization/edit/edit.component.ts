import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { VisualizationService } from '../../../services/Visualization.service';
import { SubBaseComponent } from '../../Visualization/sub.base.component';


@Component({
    selector: 'app-edit-visualization',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditVisualizationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Visualization';

    visualizationForm: FormGroup;
    visualization: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: VisualizationService,
        private fb: FormBuilder
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

    
    updateVisualization(title, options, Dashboard, Report, Metrics, Dimensions, Datasets, ChartType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateVisualization(title, options, Dashboard, Report, Metrics, Dimensions, Datasets, ChartType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexVisualization']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getVisualization(params['id']).subscribe(res => {
                this.visualization = res;
            });
        });
    }
}