import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CabinLayoutService } from '../../../services/CabinLayout.service';
import { CabinLayout } from '../../../models/CabinLayout';
import { SubBaseComponent } from '../../CabinLayout/sub.base.component';

@Component({
    selector: 'app-create-cabinLayout',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCabinLayoutComponent extends SubBaseComponent implements OnInit {

    title = 'Add CabinLayout';

    cabinLayoutForm: FormGroup;
    cabinLayout: CabinLayout;

    constructor( http: HttpClient,
        private cabinLayoutService: CabinLayoutService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.cabinLayoutForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  layoutCode: ['', Validators.required],
      totalSeats: ['', Validators.required],
      classLayout: ['', Validators.required],
      Variant: ['', ],
      Aircraft: ['', ],
      Options: ['', ]
        });
    }

    
    addCabinLayout(layoutCode, totalSeats, classLayout, Variant, Aircraft, Options): void {
        this.cabinLayoutService
        .addCabinLayout(layoutCode, totalSeats, classLayout, Variant, Aircraft, Options)
            .subscribe(() => {
                this.router.navigate(['/indexCabinLayout']);
            });
    }

    ngOnInit(): void {
    }
}