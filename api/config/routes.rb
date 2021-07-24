Rails.application.routes.draw do
  namespace :api do
    mount_devise_token_auth_for 'User', at: 'auth' , controllers: {
      registrations: 'api/auth/registrations'
    }
  end
  namespace 'api' do
    namespace 'v1' do
      resources :users, only: [:index]
    end
  end
  resources :answers
  resources :questions
  resources :logs
  resources :trees
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end
