
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexIncidentComponent } from './index.component';
import { IncidentService } from '../../../services/Incident.service';

describe('IndexIncidentComponent', () => {
  let component: IndexIncidentComponent;
  let fixture: ComponentFixture<IndexIncidentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexIncidentComponent
      ],
      providers: [
        IncidentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexIncidentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});