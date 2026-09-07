
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftPackageComponent } from './index.component';
import { AircraftPackageService } from '../../../services/AircraftPackage.service';

describe('IndexAircraftPackageComponent', () => {
  let component: IndexAircraftPackageComponent;
  let fixture: ComponentFixture<IndexAircraftPackageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftPackageComponent
      ],
      providers: [
        AircraftPackageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftPackageComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});