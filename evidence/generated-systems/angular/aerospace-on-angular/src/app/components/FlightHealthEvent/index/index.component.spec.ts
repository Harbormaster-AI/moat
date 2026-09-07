
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFlightHealthEventComponent } from './index.component';
import { FlightHealthEventService } from '../../../services/FlightHealthEvent.service';

describe('IndexFlightHealthEventComponent', () => {
  let component: IndexFlightHealthEventComponent;
  let fixture: ComponentFixture<IndexFlightHealthEventComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFlightHealthEventComponent
      ],
      providers: [
        FlightHealthEventService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFlightHealthEventComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});