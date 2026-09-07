
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexFacilityComponent } from './index.component';
import { FacilityService } from '../../../services/Facility.service';

describe('IndexFacilityComponent', () => {
  let component: IndexFacilityComponent;
  let fixture: ComponentFixture<IndexFacilityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexFacilityComponent
      ],
      providers: [
        FacilityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexFacilityComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});