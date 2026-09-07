import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RoutingService } from '../../../services/Routing.service';
import { Routing } from '../../../models/Routing';
import { SubBaseComponent } from '../../Routing/sub.base.component';

@Component({
    selector: 'app-create-routing',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRoutingComponent extends SubBaseComponent implements OnInit {

    title = 'Add Routing';

    routingForm: FormGroup;
    routing: Routing;

    constructor( http: HttpClient,
        private routingService: RoutingService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.routingForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  routingNumber: ['', Validators.required],
      revision: ['', Validators.required],
      effectivityStart: ['', Validators.required],
      effectivityEnd: ['', Validators.required],
      Item: ['', ],
      Operations: ['', ],
      RoutingType: ['', ],
      Status: ['', ]
        });
    }

    
    addRouting(routingNumber, revision, effectivityStart, effectivityEnd, Item, Operations, RoutingType, Status): void {
        this.routingService
        .addRouting(routingNumber, revision, effectivityStart, effectivityEnd, Item, Operations, RoutingType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexRouting']);
            });
    }

    ngOnInit(): void {
    }
}