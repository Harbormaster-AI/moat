import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RunParameterService } from '../../../services/RunParameter.service';
import { RunParameter } from '../../../models/RunParameter';
import { SubBaseComponent } from '../../RunParameter/sub.base.component';

@Component({
    selector: 'app-create-runParameter',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRunParameterComponent extends SubBaseComponent implements OnInit {

    title = 'Add RunParameter';

    runParameterForm: FormGroup;
    runParameter: RunParameter;

    constructor( http: HttpClient,
        private runParameterService: RunParameterService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRunParameter(name, value, TrainingRun): void {
        this.runParameterService
        .addRunParameter(name, value, TrainingRun)
            .subscribe(() => {
                this.router.navigate(['/indexRunParameter']);
            });
    }

    ngOnInit(): void {
    }
}