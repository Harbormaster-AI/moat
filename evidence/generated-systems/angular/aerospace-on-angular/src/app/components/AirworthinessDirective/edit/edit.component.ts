import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AirworthinessDirectiveService } from '../../../services/AirworthinessDirective.service';
import { SubBaseComponent } from '../../AirworthinessDirective/sub.base.component';


@Component({
    selector: 'app-edit-airworthinessDirective',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAirworthinessDirectiveComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AirworthinessDirective';

    airworthinessDirectiveForm: FormGroup;
    airworthinessDirective: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AirworthinessDirectiveService,
        private fb: FormBuilder
) {
        super(http);
        this.airworthinessDirectiveForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  directiveNumber: ['', Validators.required],
      title: ['', Validators.required],
      WorkOrders: ['', ]
        });
    }

    
    updateAirworthinessDirective(directiveNumber, title, WorkOrders): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAirworthinessDirective(directiveNumber, title, WorkOrders, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAirworthinessDirective']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAirworthinessDirective(params['id']).subscribe(res => {
                this.airworthinessDirective = res;
            });
        });
    }
}