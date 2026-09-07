import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RunParameterService } from '../../../services/RunParameter.service';
import { SubBaseComponent } from '../../RunParameter/sub.base.component';


@Component({
    selector: 'app-edit-runParameter',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRunParameterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit RunParameter';

    runParameterForm: FormGroup;
    runParameter: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RunParameterService,
        private fb: FormBuilder
) {
        super(http);
        this.runParameterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      value: ['', Validators.required],
      TrainingRun: ['', ]
        });
    }

    
    updateRunParameter(name, value, TrainingRun): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRunParameter(name, value, TrainingRun, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRunParameter']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRunParameter(params['id']).subscribe(res => {
                this.runParameter = res;
            });
        });
    }
}