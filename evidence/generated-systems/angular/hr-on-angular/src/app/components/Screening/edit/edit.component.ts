import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ScreeningService } from '../../../services/Screening.service';
import { SubBaseComponent } from '../../Screening/sub.base.component';


@Component({
    selector: 'app-edit-screening',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditScreeningComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Screening';

    screeningForm: FormGroup;
    screening: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ScreeningService,
        private fb: FormBuilder
) {
        super(http);
        this.screeningForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      completedDate: ['', Validators.required],
      Application: ['', ],
      Status: ['', ]
        });
    }

    
    updateScreening(name, completedDate, Application, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateScreening(name, completedDate, Application, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexScreening']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getScreening(params['id']).subscribe(res => {
                this.screening = res;
            });
        });
    }
}