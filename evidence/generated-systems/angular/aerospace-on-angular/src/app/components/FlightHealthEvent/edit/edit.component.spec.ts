
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditFlightHealthEventComponent } from './edit.component';
import { FlightHealthEventService } from '../../../services/FlightHealthEvent.service';

describe('EditFlightHealthEventComponent', () => {
  let component: EditFlightHealthEventComponent;
  let fixture: ComponentFixture<EditFlightHealthEventComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditFlightHealthEventComponent
      ],
      providers: [
        FlightHealthEventService,
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

    fixture = TestBed.createComponent(EditFlightHealthEventComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});