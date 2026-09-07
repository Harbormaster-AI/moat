
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInboundShipmentLineComponent } from './create.component';
import { InboundShipmentLineService } from '../../../services/InboundShipmentLine.service';
import { Router } from '@angular/router';

describe('CreateInboundShipmentLineComponent', () => {
  let component: CreateInboundShipmentLineComponent;
  let fixture: ComponentFixture<CreateInboundShipmentLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInboundShipmentLineComponent
      ],
      providers: [
        InboundShipmentLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInboundShipmentLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});