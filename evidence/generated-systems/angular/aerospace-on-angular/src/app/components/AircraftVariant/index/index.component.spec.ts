
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftVariantComponent } from './index.component';
import { AircraftVariantService } from '../../../services/AircraftVariant.service';

describe('IndexAircraftVariantComponent', () => {
  let component: IndexAircraftVariantComponent;
  let fixture: ComponentFixture<IndexAircraftVariantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftVariantComponent
      ],
      providers: [
        AircraftVariantService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftVariantComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});