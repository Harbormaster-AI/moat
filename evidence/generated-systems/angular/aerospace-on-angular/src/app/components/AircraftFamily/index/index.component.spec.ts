
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftFamilyComponent } from './index.component';
import { AircraftFamilyService } from '../../../services/AircraftFamily.service';

describe('IndexAircraftFamilyComponent', () => {
  let component: IndexAircraftFamilyComponent;
  let fixture: ComponentFixture<IndexAircraftFamilyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftFamilyComponent
      ],
      providers: [
        AircraftFamilyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftFamilyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});