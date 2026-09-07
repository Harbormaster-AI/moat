import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ExperimentVariantService } from '../../../services/ExperimentVariant.service';
import { SubBaseComponent } from '../../ExperimentVariant/sub.base.component';


@Component({
    selector: 'app-edit-experimentVariant',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditExperimentVariantComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ExperimentVariant';

    experimentVariantForm: FormGroup;
    experimentVariant: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ExperimentVariantService,
        private fb: FormBuilder
) {
        super(http);
        this.experimentVariantForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      allocation: ['', Validators.required],
      Experiment: ['', ],
      CreativeVariation: ['', ],
      LineItem: ['', ]
        });
    }

    
    updateExperimentVariant(name, allocation, Experiment, CreativeVariation, LineItem): void {
        this.route.params.subscribe((params) => {

                        this.service.updateExperimentVariant(name, allocation, Experiment, CreativeVariation, LineItem, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexExperimentVariant']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getExperimentVariant(params['id']).subscribe(res => {
                this.experimentVariant = res;
            });
        });
    }
}