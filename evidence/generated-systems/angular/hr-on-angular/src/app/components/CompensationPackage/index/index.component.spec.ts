
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCompensationPackageComponent } from './index.component';
import { CompensationPackageService } from '../../../services/CompensationPackage.service';

describe('IndexCompensationPackageComponent', () => {
  let component: IndexCompensationPackageComponent;
  let fixture: ComponentFixture<IndexCompensationPackageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCompensationPackageComponent
      ],
      providers: [
        CompensationPackageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCompensationPackageComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});