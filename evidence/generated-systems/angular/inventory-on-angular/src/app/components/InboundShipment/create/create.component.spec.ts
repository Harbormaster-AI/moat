
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInboundShipmentComponent } from './create.component';
import { InboundShipmentService } from '../../../services/InboundShipment.service';
import { Router } from '@angular/router';

describe('CreateInboundShipmentComponent', () => {
  let component: CreateInboundShipmentComponent;
  let fixture: ComponentFixture<CreateInboundShipmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInboundShipmentComponent
      ],
      providers: [
        InboundShipmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInboundShipmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});