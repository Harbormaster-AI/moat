
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexServiceProviderComponent } from './index.component';
import { ServiceProviderService } from '../../../services/ServiceProvider.service';

describe('IndexServiceProviderComponent', () => {
  let component: IndexServiceProviderComponent;
  let fixture: ComponentFixture<IndexServiceProviderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexServiceProviderComponent
      ],
      providers: [
        ServiceProviderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexServiceProviderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});