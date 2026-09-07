
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditInboundShipmentLineComponent } from './edit.component';
import { InboundShipmentLineService } from '../../../services/InboundShipmentLine.service';

describe('EditInboundShipmentLineComponent', () => {
  let component: EditInboundShipmentLineComponent;
  let fixture: ComponentFixture<EditInboundShipmentLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditInboundShipmentLineComponent
      ],
      providers: [
        InboundShipmentLineService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditInboundShipmentLineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});