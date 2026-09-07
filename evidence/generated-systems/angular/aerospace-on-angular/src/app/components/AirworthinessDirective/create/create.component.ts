import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AirworthinessDirectiveService } from '../../../services/AirworthinessDirective.service';
import { AirworthinessDirective } from '../../../models/AirworthinessDirective';
import { SubBaseComponent } from '../../AirworthinessDirective/sub.base.component';

@Component({
    selector: 'app-create-airworthinessDirective',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAirworthinessDirectiveComponent extends SubBaseComponent implements OnInit {

    title = 'Add AirworthinessDirective';

    airworthinessDirectiveForm: FormGroup;
    airworthinessDirective: AirworthinessDirective;

    constructor( http: HttpClient,
        private airworthinessDirectiveService: AirworthinessDirectiveService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAirworthinessDirective(directiveNumber, title, WorkOrders): void {
        this.airworthinessDirectiveService
        .addAirworthinessDirective(directiveNumber, title, WorkOrders)
            .subscribe(() => {
                this.router.navigate(['/indexAirworthinessDirective']);
            });
    }

    ngOnInit(): void {
    }
}