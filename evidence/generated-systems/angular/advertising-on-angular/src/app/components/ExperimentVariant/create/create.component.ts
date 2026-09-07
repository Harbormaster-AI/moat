import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ExperimentVariantService } from '../../../services/ExperimentVariant.service';
import { ExperimentVariant } from '../../../models/ExperimentVariant';
import { SubBaseComponent } from '../../ExperimentVariant/sub.base.component';

@Component({
    selector: 'app-create-experimentVariant',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateExperimentVariantComponent extends SubBaseComponent implements OnInit {

    title = 'Add ExperimentVariant';

    experimentVariantForm: FormGroup;
    experimentVariant: ExperimentVariant;

    constructor( http: HttpClient,
        private experimentVariantService: ExperimentVariantService,
        private fb: FormBuilder,
        private router: Router
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

    
    addExperimentVariant(name, allocation, Experiment, CreativeVariation, LineItem): void {
        this.experimentVariantService
        .addExperimentVariant(name, allocation, Experiment, CreativeVariation, LineItem)
            .subscribe(() => {
                this.router.navigate(['/indexExperimentVariant']);
            });
    }

    ngOnInit(): void {
    }
}