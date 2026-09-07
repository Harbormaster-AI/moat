
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInboundShipmentLineComponent } from './index.component';
import { InboundShipmentLineService } from '../../../services/InboundShipmentLine.service';

describe('IndexInboundShipmentLineComponent', () => {
  let component: IndexInboundShipmentLineComponent;
  let fixture: ComponentFixture<IndexInboundShipmentLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInboundShipmentLineComponent
      ],
      providers: [
        InboundShipmentLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInboundShipmentLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});