
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInboundShipmentComponent } from './index.component';
import { InboundShipmentService } from '../../../services/InboundShipment.service';

describe('IndexInboundShipmentComponent', () => {
  let component: IndexInboundShipmentComponent;
  let fixture: ComponentFixture<IndexInboundShipmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInboundShipmentComponent
      ],
      providers: [
        InboundShipmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInboundShipmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});