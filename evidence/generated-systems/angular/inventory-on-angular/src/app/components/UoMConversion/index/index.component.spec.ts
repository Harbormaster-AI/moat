
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexUoMConversionComponent } from './index.component';
import { UoMConversionService } from '../../../services/UoMConversion.service';

describe('IndexUoMConversionComponent', () => {
  let component: IndexUoMConversionComponent;
  let fixture: ComponentFixture<IndexUoMConversionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexUoMConversionComponent
      ],
      providers: [
        UoMConversionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexUoMConversionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});