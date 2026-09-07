import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PositionService } from '../../../services/Position.service';
import { SubBaseComponent } from '../../Position/sub.base.component';


@Component({
    selector: 'app-edit-position',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPositionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Position';

    positionForm: FormGroup;
    position: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PositionService,
        private fb: FormBuilder
) {
        super(http);
        this.positionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  quantity: ['', Validators.required],
      averageCost: ['', Validators.required],
      marketValue: ['', Validators.required],
      Portfolio: ['', ],
      Security: ['', ]
        });
    }

    
    updatePosition(quantity, averageCost, marketValue, Portfolio, Security): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePosition(quantity, averageCost, marketValue, Portfolio, Security, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPosition']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPosition(params['id']).subscribe(res => {
                this.position = res;
            });
        });
    }
}